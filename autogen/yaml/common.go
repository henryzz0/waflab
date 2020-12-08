package yaml

//
//type YAML struct {
//	Meta  Meta `yaml:"meta"`
//	Tests Tests `yaml:"tests"`
//}
//
//type Meta struct {
//	Author      string `yaml:"author"`
//	Description string `yaml:"description"`
//	Enabled     bool   `yaml:"enabled"`
//	Name        string `yaml:"name"`
//}
//
//type Tests struct {
//	Stages []Stages `yaml:"stages"`
//}
//
//type Stages struct {
//	Stage []Stage `yaml:"stage"`
//	TestTitle string `yaml:"test_title"`
//}
//
//type Stage struct {
//	Input  Input  `yaml:"input"`
//	Output Output `yaml:"output"`
//}
//
//type Input struct {
//	DestAddr       string            `yaml:"dest_addr"`
//	Port           int               `yaml:"port"`
//	Method         string            `yaml:"method"`
//	Headers        map[string]string `yaml:"headers,omitempty"`
//	Protocol       string            `yaml:"protocol"`
//	Uri            string            `yaml:"uri"`
//	Version        string            `yaml:"version"`
//	Data           string            `yaml:"data,omitempty"`
//	SaveCookie     bool              `yaml:"save_cookie,omitempty"`
//	StopMagic      bool              `yaml:"stop_magic,omitempty"`
//	EncodedRequest string            `yaml:"encoded_request,omitempty"`
//	RawRequest     string            `yaml:"raw_request,omitempty"`
//}
//
//type Output struct {
//	Status int `yaml:"status"`
//}
