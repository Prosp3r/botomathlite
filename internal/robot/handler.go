package robot

type Robot struct {
	ID        int64        `json:"id"`
	RobotName string       `json:"robotname"`
	RoboType  string       `json:"robotype"`
	Skill     []RobotSkill `json:"skill"`
}

type RobotSkill struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score string `json:"score"`
}
