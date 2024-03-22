package data

import "github.com/Piszmog/website/models"

var ExperienceData = []models.TimelineEntry{
	models.Job{
		Start:        "2023-11-01",
		End:          "Present",
		Company:      "HelloFresh",
		Title:        "Senior Software Engineer",
		ChangeReason: models.ChangeReasonNew,
		Toolbox: models.Toolbox{
			Plan:    []string{"JIRA"},
			Code:    []string{"Go", "Kotlin", "TypeScript"},
			Model:   []string{"JSON", "Protobuf", "GraphQL"},
			View:    []string{"React"},
			Build:   []string{"GitHub Actions"},
			Run:     []string{"AWS", "Kubernetes", "Docker"},
			Persist: []string{"PostgreSQL"},
			Move:    []string{"GraphQL"},
		},
		Details: []string{
			"Enhanced and maintained a suite of Kotlin-based microservices, instrumental in optimizing the company’s ready-to-eat product manufacturing and planning processes.",
		},
	},
	models.Job{
		Start:        "2022-05-01",
		End:          "2023-08-01",
		Company:      "Sourcegraph",
		Title:        "Senior Software Engineer",
		ChangeReason: models.ChangeReasonNew,
		Toolbox: models.Toolbox{
			Plan:    []string{"GitHub"},
			Code:    []string{"Go", "TypeScript", "Python", "JavaScript"},
			Model:   []string{"JSON", "Protobuf", "GraphQL"},
			View:    []string{"React", "Svelte"},
			Build:   []string{"GitHub Actions", "Buildkite"},
			Run:     []string{"Google Cloud Platform", "AWS", "Kubernetes", "Docker"},
			Persist: []string{"PostgreSQL", "Redis"},
			Move:    []string{"REST", "gRPC", "GraphQL"},
		},
		Details: []string{
			"Led the design and deployment of advanced search capabilities, enabling security teams and auditors to conduct in-depth analysis of Git codebases.",
			"Directed a team of 3 developers to create a remote code execution service, enabling users to run run arbitrary code within a secured environment with Kubernetes and Firecracker.",
			"Built long asked for features into the Batch Changes product, using Go and React.",
		},
	},
	models.Job{
		Start:        "2019-07-01",
		End:          "2022-05-01",
		Company:      "Charles Schwab",
		Title:        "Senior Software Engineer",
		ChangeReason: models.ChangeReasonPromotion,
		Toolbox: models.Toolbox{
			Plan:    []string{"JIRA", "GitHub"},
			Code:    []string{"Java", "Spring Boot", "Go", "TypeScript", "Python", "JavaScript"},
			Model:   []string{"JSON"},
			View:    []string{"Angular"},
			Build:   []string{"Bamboo", "GitHub Actions"},
			Run:     []string{"Cloud Foundry"},
			Persist: []string{"Microsoft SQL Server"},
			Move:    []string{"REST"},
		},
		Details: []string{
			"Drove the creation of a proxy sidecar service written in Go, enhancing Enterprise services with standardized, logging, advanced security measures, and seamless service discovery.",
			"Led a team of 3 developers in creating an InnerSource platform that empowered enterprise teams to share code, enabled management to assess risks, and provided leadership with tools to monitor adoption.",
			"Built GitHub Actions, in JavaScript, to seamlessly onboard and transition Enterprise teams to GitHub.",
			"Developed a suite of InnerSource libraries for streamlined logging, automated deployments, and service discovery across .NET, Java, and Python applications.",
		},
	},
	models.Job{
		Start:        "2017-03-01",
		End:          "2019-07-01",
		Company:      "Charles Schwab",
		Title:        "Software Engineer",
		ChangeReason: models.ChangeReasonPromotion,
		Toolbox: models.Toolbox{
			Plan:    []string{"JIRA"},
			Code:    []string{"Java", "Spring Boot", "Python", "JavaScript"},
			Model:   []string{"XML", "JSON"},
			View:    []string{"Angular"},
			Build:   []string{"Bamboo"},
			Run:     []string{"Cloud Foundry"},
			Persist: []string{"MongoDB", "Redis"},
			Move:    []string{"REST", "SOAP", "RabbitMQ"},
		},
		Details: []string{
			"Designed and implemented 30+ microservices for efficient non-marketing communications, handling millions of records daily.",
			"Led the transition from Mainframe processes to a contemporary platform leveraging Cloud technologies and Java.",
			"Automated our deployment ticketing process using Python, greatly decreasing the amount of time a team member needed to spend on creating tickets.",
		},
	},
	models.Job{
		Start:        "2016-01-01",
		End:          "2017-03-01",
		Company:      "Charles Schwab",
		Title:        "Associate Software Engineer",
		ChangeReason: models.ChangeReasonNew,
		Toolbox: models.Toolbox{
			Plan:    []string{"Team Foundation Server", "JIRA"},
			Code:    []string{"Java", "Spring"},
			Model:   []string{"XML", "JSON"},
			View:    []string{},
			Build:   []string{"Jenkins", "Bamboo"},
			Run:     []string{"WebLogic", "Cloud Foundry"},
			Persist: []string{"MongoDB"},
			Move:    []string{"REST", "SOAP"},
		},
		Details: []string{
			"Provided support for legacy Java Spring applications focused on non-marketing communications processing.",
			"Enhanced application reliability by elevating test coverage and enriching documentation for better clarity and usability.",
			"Led my team in being among the trailblazers to deploy applications on Pivotal Cloud Foundry.",
			"Steered my team's seamless transition from Team Foundation Server to the Atlassian suite, emphasizing our commitment to Agile methodologies.",
		},
	},
	models.Job{
		Start:        "2015-03-01",
		End:          "2016-01-01",
		Company:      "Raytheon",
		Title:        "Software Engineer",
		ChangeReason: models.ChangeReasonNew,
		Toolbox: models.Toolbox{
			Plan:    []string{},
			Code:    []string{"Java", "Swing", "JavaScript", "Python"},
			Model:   []string{"XML", "JSON"},
			View:    []string{"PrimeFaces"},
			Build:   []string{"Ant", "Python"},
			Run:     []string{"WebLogic"},
			Persist: []string{"Oracle"},
			Move:    []string{"REST", "SOAP"},
		},
		Details: []string{
			"Led the performance of satellite management applications through advanced algorithm optimization and augmented the Thick Client with innovative features.",
			"Built a lightweight web application alternative, catering to users operating on constrained hardware, offering an efficient replacement to the Thick Client.",
			"Led my team's shift from Ant to Python for application development, significantly reducing build time.",
		},
	},
	models.Job{
		Start:        "2014-01-01",
		End:          "2015-03-01",
		Company:      "Intelligent Software Solutions",
		Title:        "Software Engineer",
		ChangeReason: models.ChangeReasonNew,
		Toolbox: models.Toolbox{
			Plan:    []string{},
			Code:    []string{"Java", "Swing", "JavaScript", "Spring"},
			Model:   []string{"XML", "JSON"},
			View:    []string{"Ext JS", "PrimeFaces"},
			Build:   []string{"TeamCity"},
			Run:     []string{"WebSphere"},
			Persist: []string{"Oracle"},
			Move:    []string{"REST", "SOAP", "RabbitMQ"},
		},
		Details: []string{
			"Developed and integrated enhancements for a space debris monitoring web application using HTML, JavaScript, and jQuery, significantly improving user experience.",
			"Led improvements to both web and Thick Client applications for Strategy and Assessment, trusted tools of the U.S. Air Force and the Royal Canadian Air Force.",
		},
	},
	models.Education{
		Start:  "2009-08-01",
		End:    "2013-08-01",
		School: "University of Colorado Boulder",
		Degree: "Bachelor of Science",
		Field:  "Applied Mathematics",
	},
}

var ProjectsData = []models.Project{
	{
		Name:        "cloudconfigclient",
		Description: "Go client library for Spring Config Server",
		URL:         "https://github.com/Piszmog/cloudconfigclient",
		Languages:   []string{"Go"},
	},
	{
		Name:        "lopper",
		Description: "✂️ Deletes dead local Git branches",
		URL:         "https://github.com/Piszmog/lopper",
		Languages:   []string{"Go"},
	},
	{
		Name:        "cloud-config-client-autoconfig",
		Description: "Auto Configuration for creating Config Clients",
		URL:         "https://github.com/Piszmog/cloud-config-client-autoconfig",
		Languages:   []string{"Java", "Spring Boot"},
	},
	{
		Name:        "next-version",
		Description: "Action that increments a project to the next version at pull request time.",
		URL:         "https://github.com/Piszmog/next-version",
		Languages:   []string{"JavaScript"},
	},
	{
		Name:        "cf-services",
		Description: "Retrieve configurations from Cloud Foundry with Rust.",
		URL:         "https://github.com/Piszmog/cf-services",
		Languages:   []string{"Rust"},
	},
	{
		Name:        "job-app-tracker",
		Description: "📈 Track and manage your job applications",
		URL:         "https://github.com/Piszmog/job-app-tracker",
		Languages:   []string{"Svelte", "Rust", "TypeScript"},
	},
	{
		Name:        "website",
		Description: "🌐 My personal website",
		URL:         "https://github.com/Piszmog/website",
		Languages:   []string{"Go", "HTML", "CSS"},
	},
	{
		Name:        "jiggle",
		Description: "🧂 Merges changes through a git tree.",
		URL:         "https://github.com/Piszmog/jiggle",
		Languages:   []string{"Rust"},
	},
	{
		Name:        "pathwise",
		Description: "📈 Track and manage your job applications",
		URL:         "https://github.com/Piszmog/pathwise",
		Languages:   []string{"Go", "CSS"},
	},
}
