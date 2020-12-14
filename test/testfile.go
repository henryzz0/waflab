package test

type Testfile struct {
	Meta  *Meta   `yaml:"meta"`
	Tests []*Test `yaml:"tests"`
}

type Meta struct {
	Author      string `yaml:"author"`
	Enabled     bool   `yaml:"enabled"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Test struct {
	TestTitle string          `yaml:"test_title"`
	Desc      string          `yaml:"desc"`
	Stages    []*StageWrapper `yaml:"stages"`
}

type StageWrapper struct {
	Stage *Stage `yaml:"stage"`
}

type Stage struct {
	Input  *Input  `yaml:"input"`
	Output *Output `yaml:"output"`
}

type Input struct {
	SaveCookie     bool              `yaml:"save_cookie,omitempty"`
	DestAddr       string            `yaml:"dest_addr"`
	Method         string            `yaml:"method"`
	Port           int               `yaml:"port"`
	Protocol       string            `yaml:"protocol"`
	Uri            string            `yaml:"uri"`
	Version        string            `yaml:"version"`
	Headers        map[string]string `yaml:"headers,omitempty"`
	Data           []string          `yaml:"data,omitempty"`
	EncodedRequest string            `yaml:"encoded_request,omitempty"`
	RawRequest     string            `yaml:"raw_request,omitempty"`
}

type Output struct {
	Status        []int  `yaml:"status"`
	HtmlContains  string `yaml:"html_contains,omitempty"`
	LogContains   string `yaml:"log_contains,omitempty"`
	NoLogContains string `yaml:"no_log_contains,omitempty"`
}
