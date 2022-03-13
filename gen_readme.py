# -*- coding: UTF-8 -*-
import os
import requests
import json
import re
import sys

# usage: python leetcode.py <username> <dir> <num>
# > Please input topic link:

Remove = ["</?p>", "</?ul>", "</?ol>", "</li>", "</sup>"]
Replace = [["&nbsp;", " "], ["&quot;", '"'], ["&lt;", "<"], ["&gt;", ">"],
           ["&le;", "≤"], ["&ge;", "≥"], ["<sup>", "^"], ["&#39", "'"],
           ["&times;", "x"], ["&ldquo;", "“"], ["&rdquo;", "”"],
           [" *<strong> *", " **"], [" *</strong> *", "** "],
           [" *<code> *", " `"], [" *</code> *", "` "], ["<pre>", "```\n"],
           ["</pre>", "\n```\n"], ["<em> *</em>", ""], [" *<em> *", " *"],
           [" *</em> *", "* "], ["</?div.*?>", ""], ["	*</?li>", "- "]]

markdown_format = """# {titlename}

## 1. 题目描述

{Content}

## 2. 解题

"""

graphql_query = '''query getQuestionDetail($titleSlug: String!) {
            question(titleSlug: $titleSlug) {
                questionId
                questionFrontendId
                title
                titleSlug
                content
                translatedTitle
                translatedContent
                difficulty
                topicTags {
                    name
                    slug
                    translatedName
                    __typename
                }
                codeSnippets {
                    lang
                    langSlug
                    code
                    __typename
                }
                __typename
            }
        }'''

gotest_template = """import (
    "testing"
)

func Test_{0}(t *testing.T){{
    panic("not implemented") // TODO: Implement
}}
"""


def convert(src):
    # pre内部预处理
    def remove_label_in_pre(matched):
        tmp = matched.group()
        tmp = re.sub("<[^>p]*>", "", tmp)  # 不匹配>与p
        return tmp

    src = re.sub("<pre>[\s\S]*?</pre>", remove_label_in_pre,
                 src)  # 注意此处非贪心匹配，因为可能有多个示例
    # 可以直接删除的标签
    for curPattern in Remove:
        src = re.sub(curPattern, "", src)
    # 需要替换内容的标签
    for curPattern, curRepl in Replace:
        src = re.sub(curPattern, curRepl, src)
    return src


def get_all(slug: str) -> dict:
    """
    根据题目的 id 获取题目的名字，内容，代码块，标签，返回 json 格式的 dict 对象
    """
    user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36"
    session = requests.Session()
    url = "https://leetcode-cn.com/graphql"
    params = {
        'operationName':
        "getQuestionDetail",
        'variables': {
            'titleSlug': slug
        },
        'query': graphql_query
    }
    json_data = json.dumps(params).encode('utf8')
    headers = {
        'User-Agent': user_agent,
        'Connection': 'keep-alive',
        'Content-Type': 'application/json',
        'Referer': 'https://leetcode-cn.com/problems/' + slug
    }
    resp = session.post(url, data=json_data, headers=headers, timeout=10)
    resp.encoding = 'utf8'
    content = resp.json()
    question = content['data']['question']
    return question


def get_problem_content(slug: str) -> str:
    question = get_all(slug)
    res = convert(question['translatedContent'])
    # 在正文后面填上标签
    res += "\n \n**标签**\n"
    tags = question['topicTags']
    for tag in tags:
        if tag['translatedName'] != None:
            tagName = tag['translatedName']
        else:
            tagName = tag['name']
        res += "`" + tagName + "` "

    return res + "\n"


def get_solution_by_lang(slug: str, lang: str) -> str:
    """
        获取给定题目的对应语言的函数

        支持的参数如下

        C++ Java Python Python3 C C# JavaScript Ruby Swift Go Scala Kotlin
        Rust PHP TypeScript Racket
    """
    question = get_all(slug)
    # 获取对应语言的函数
    codeSnippets = question['codeSnippets']
    for x in codeSnippets:
        if x['lang'] == lang:
            return x['code']


def gen_markdown(path, content, title):
    """
        生成 markdown 文件
    """
    file = open(path, 'a', encoding="utf-8")
    markdown_content = markdown_format.format(
        titlename=title, Content=content)
    file.write(markdown_content)
    print("Generate readme.md in: ", path)
    file.close()


def gen_files(url: str, username: str, dir: str, num: str):
    """
        根据 url 生成对应的文件
    """

    if url.startswith("https://leetcode-cn.com/problems/"):
        slug = url.replace("https://leetcode-cn.com/problems/", "",
                           1).strip('/')
    else:
        print("Generate readme.md failed. Got Wrong URL!!!\nIt should be like https://leetcode-cn.com/problems/evaluate-division/")
        return

    url = "https://leetcode-cn.com/problems/" + slug

    # 获取一些变量
    question = get_all(slug=slug)
    title = question['questionFrontendId'] + '.' + question['translatedTitle']

    content = get_problem_content(slug)

    raw_func = get_solution_by_lang(slug, 'Go')
    func = re.compile(r'({\n.*?\n})').sub(
        "{\n\tpanic(\"not implemented\") // TODO: Implement\n}", raw_func, 1)
    content = re.sub(r'\n\n\n\n*', "\n", content)  # 替换掉多个换行符

    func_name = re.compile(r'(func\s\w+)').findall(raw_func)[0].split(" ")[-1]
    test_func = gotest_template.format(func_name)

    # 获取当前目录
    base_dir = os.getcwd()

    # 生成 markdown 文件
    go_package = dir+num
    filepath = os.path.join(base_dir, username,
                            dir, go_package, "readme.md")
    gen_markdown(filepath, content, title)

    # 向go文件中写入解题函数
    with open(os.path.join(base_dir, username, dir, go_package, go_package+".go"), mode='a') as go_file:
        go_file.write("\n")
        go_file.write(func)

    # 向go_test文件中写入测试函数
    with open(os.path.join(base_dir, username, dir, go_package, go_package+"_test"+".go"), mode='a') as test_file:
        test_file.write("\n")
        test_file.write(test_func)
    return


if __name__ == '__main__':
    url = input("请输入题目链接: ")
    gen_files(url=url, username=sys.argv[1], dir=sys.argv[2], num=sys.argv[3])
    # gen_files(url=url, username="yangchnet", dir="goden", num="0103")
