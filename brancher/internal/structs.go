package internal

type livenessProbeResponse struct {
	Code int 			`json:"code" xml:"code"`
	Response string 	`json:"response" xml:"response"`
}

type readinessProbeResponse struct {
	Code int			`json:"code" xml:"code"`
	Response string		`json:"response" xml:"response"`
	Postgres string		`json:"postgres" xml:"postgres"`
	Rabbitmq string		`json:"rabbitmq" xml:"rabbitmq"`
}