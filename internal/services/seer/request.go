package seer

type Request struct {
	ProgrammingLanguages ProgrammingLanguages `json:"programming_languages"`
	BackendFrameworks    BackendFrameworks    `json:"backend_frameworks"`
	Devops               Devops               `json:"devops"`
	FrontendFrameworks   FrontendFrameworks   `json:"frontend_frameworks"`
	SystemDesignSkills   SystemDesignSkills   `json:"system_design_skills"`
	VersionControls      VersionControls      `json:"version_controls"`
	UIUX                 UIUX                 `json:"ui_ux"`
	DataRelated          DataRelated          `json:"data_related"`
	Databases            Databases            `json:"databases"`
	AIModeling           AIModeling           `json:"ai_modeling"`
	AIFrameworks         AIFrameworks         `json:"ai_frameworks"`
	NetworkSkills        NetworkSkills        `json:"network_skills"`
	HardwareSkills       HardwareSkills       `json:"hardware_skills"`
	Education            Education            `json:"education"`
	Experience           Experience           `json:"experience"`
}

type ProgrammingLanguages struct {
	Java       int `json:"java"`
	Kotlin     int `json:"kotlin"`
	XML        int `json:"xml"`
	Swift      int `json:"swift"`
	R          int `json:"r"`
	Python     int `json:"python"`
	PHP        int `json:"php"`
	Javascript int `json:"javascript"`
	HTML       int `json:"html"`
	CSS        int `json:"css"`
	Go         int `json:"go"`
	Cpp        int `json:"cpp"`
	C          int `json:"c"`
	Csharp     int `json:"csharp"`
	Ruby       int `json:"ruby"`
	Perl       int `json:"perl"`
	Bash       int `json:"bash"`
	Scala      int `json:"scala"`
}

type BackendFrameworks struct {
	BackendDevelopment int `json:"backend_development"`
	ASP                int `json:"asp"`
	Django             int `json:"django"`
	Spring             int `json:"spring"`
	Laravel            int `json:"laravel"`
	Node               int `json:"node"`
	Express            int `json:"express"`
}

type Devops struct {
	Devops              int `json:"devops"`
	Kubernetes          int `json:"kubernetes"`
	Docker              int `json:"docker"`
	Ansible             int `json:"ansible"`
	Infrastructure      int `json:"infrastructure"`
	GoogleCloudPlatform int `json:"google_cloud_platform"`
	Azure               int `json:"azure"`
	Containerization    int `json:"containerization"`
	AWS                 int `json:"aws"`
	Jenkins             int `json:"jenkins"`
	CICD                int `json:"ci_cd"`
	Cloud               int `json:"cloud"`
}

type FrontendFrameworks struct {
	React     int `json:"react"`
	Vue       int `json:"vue"`
	Angular   int `json:"angular"`
	Redux     int `json:"redux"`
	Ajax      int `json:"ajax"`
	BootStrap int `json:"bootstrap"`
}

type SystemDesignSkills struct {
	MobileApplication   int `json:"mobile_application"`
	OOP                 int `json:"oop"`
	DesignPatterns      int `json:"design_patterns"`
	SoftwareEngineering int `json:"software_engineering"`
	Programming         int `json:"programming"`
	Algorithm           int `json:"algorithm"`
	WebDevelopment      int `json:"web_development"`
	WebApplications     int `json:"web_applications"`
	Microservices       int `json:"microservices"`
	SoftwareTesting     int `json:"software_testing"`
	UML                 int `json:"uml"`
}

type VersionControls struct {
	Git int `json:"git"`
}

type UIUX struct {
	GraphicDesign int `json:"graphic_design"`
	Illustrator   int `json:"illustrator"`
	Figma         int `json:"figma"`
	Photoshop     int `json:"photoshop"`
	UI_UX         int `json:"ui_ux"`
}

type DataRelated struct {
	DataIntegration      int `json:"data_integration"`
	BigData              int `json:"big_data"`
	Data                 int `json:"data"`
	Spark                int `json:"spark"`
	InformationRetrieval int `json:"information_retrieval"`
	DataVisualization    int `json:"data_visualization"`
	Selenium             int `json:"selenium"`
	PySpark              int `json:"pyspark"`
	Mapreduce            int `json:"mapreduce"`
	Zeppelin             int `json:"zeppelin"`
	Hadoop               int `json:"hadoop"`
	Sqoop                int `json:"sqoop"`
	Tableau              int `json:"tableau"`
	Hdfs                 int `json:"hdfs"`
	Oozie                int `json:"oozie"`
	DataWarehousing      int `json:"data_warehousing"`
	Ambari               int `json:"ambari"`
	Impala               int `json:"impala"`
	Snowflake            int `json:"snowflake"`
}

type Databases struct {
	DataMigration  int `json:"data_migration"`
	Mongodb        int `json:"mongodb"`
	DatabaseDesign int `json:"database_design"`
	NoSQL          int `json:"no_sql"`
	MySQL          int `json:"my_sql"`
	PostgreSQL     int `json:"postgre_sql"`
	Sqlite         int `json:"sqlite"`
	GraphQL        int `json:"graph_ql"`
	ElasticSearch  int `json:"elastic_search"`
	Redis          int `json:"redis"`
	DataCenter     int `json:"data_center"`
	MariaDB        int `json:"maria_db"`
	Oracle         int `json:"oracle"`
	SQL            int `json:"sql"`
}

type AIModeling struct {
	AI                 int `json:"ai"`
	RegressionTesting  int `json:"regression_testing"`
	PredictiveModeling int `json:"predictive_modeling"`
	MachineLearning    int `json:"machine_learning"`
	ImageProcessing    int `json:"image_processing"`
	DataScience        int `json:"data_science"`
	NLP                int `json:"nlp"`
	DataModeling       int `json:"data_modeling"`
	ComputerVision     int `json:"computer_vision"`
	BigData            int `json:"big_data"`
	DeepLearning       int `json:"deep_learning"`
	DataMining         int `json:"data_mining"`
	DataAnalysis       int `json:"data_analysis"`
	MLOps              int `json:"ml_ops"`
	SupervisedLearning int `json:"supervised_learning"`
	RecommenderSystems int `json:"recommender_systems"`
	DataPreprocessing  int `json:"data_preprocessing"`
}

type AIFrameworks struct {
	Pandas      int `json:"pandas"`
	Pytorch     int `json:"pytorch"`
	Keras       int `json:"keras"`
	Tensorflow  int `json:"tensorflow"`
	ScikitLearn int `json:"scikit_learn"`
	Seaborn     int `json:"seaborn"`
}

type NetworkSkills struct {
	DHCP                  int `json:"dhcp"`
	DNS                   int `json:"dns"`
	RoutingProtocols      int `json:"routing_protocols"`
	Routers               int `json:"routers"`
	Cisco                 int `json:"cisco"`
	Wireless              int `json:"wireless"`
	VPN                   int `json:"vpn"`
	Switching             int `json:"switching"`
	NetworkSecurity       int `json:"network_security"`
	LanWan                int `json:"lan_wan"`
	PenetrationTesting    int `json:"penetration_testing"`
	NetworkAdministration int `json:"network_administration"`
	IP                    int `json:"ip"`
	TCPIP                 int `json:"tcp_ip"`
}

type HardwareSkills struct {
	FPGA                  int `json:"fpga"`
	PCB                   int `json:"pcb"`
	VLSI                  int `json:"vlsi"`
	IOT                   int `json:"iot"`
	CircuitDesign         int `json:"circuit_design"`
	Embedded              int `json:"embedded"`
	ComputerArchitecture  int `json:"computer_architecture"`
	PSpice                int `json:"pspice"`
	MicroController       int `json:"micro_controller"`
	ControllerAreaNetwork int `json:"controller_area_network"`
	SignalProcessing      int `json:"signal_processing"`
	RaspberryPi           int `json:"raspberry_pi"`
	Verilog               int `json:"verilog"`
	STM32                 int `json:"stm32"`
	ARM                   int `json:"arm"`
	Xilinx                int `json:"xilinx"`
	Arduino               int `json:"arduino"`
	Simulink              int `json:"simulink"`
	Matlab                int `json:"matlab"`
	VHDL                  int `json:"vhdl"`
	Proteus               int `json:"proteus"`
}

type Education struct {
	BachelorComputer int `json:"bachelor_computer"`
	MasterComputer   int `json:"master_computer"`
	PHDComputer      int `json:"phd_computer"`
	OtherMajor       int `json:"other_major"`
}

type Experience struct {
	DataScientist         int `json:"data_scientist"`
	UIUXDesigner          int `json:"ui_ux_designer"`
	NetworkEngineer       int `json:"network_engineer"`
	DataEngineer          int `json:"data_engineer"`
	SoftwareEngineer      int `json:"software_engineer"`
	FrontendDeveloper     int `json:"frontend_developer"`
	HardwareEngineer      int `json:"hardware_engineer"`
	DevopsEngineer        int `json:"devops_engineer"`
	DatabaseAdministrator int `json:"database_administrator"`
}
