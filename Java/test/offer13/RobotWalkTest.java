package offer13;

import org.junit.Test;
import static org.junit.Assert.*;

public class RobotWalkTest {
    @Test
    public void main(){
        RobotWalk robot1 = new RobotWalk();
        assert(robot1.movingCount(2, 3, 1)==3);

        RobotWalk robot2 = new RobotWalk();
        assert(robot2.movingCount(3, 1, 0)==1);

    }

}