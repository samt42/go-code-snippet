package main


import (
  "fmt"
  "io"
  "os"
  "strings"
  "bufio"
  
)

func main() {
  conf := setConfig("config.ini")
  ak := GetValue(conf, "AK", "AK")
  sk := GetValue(conf, "SK", "SK")
  fmt.Println("SK : ", sk) 
  fmt.Println("AK : ", ak)
}

//config结构体
type Config struct{
  filepath string
  conflist []map[string]map[string]string
}

//获取文件路径
func setConfig(filepath string) *Config{
  c := new(Config)
  c.filepath = filepath

  return c
}

//获取响应配置文件的值
func GetValue(c *Config, section string, name string) string {
  conf := ReadList(c)
  for _, v := range conf {
    for key, value := range v {
      if key == section {
        return value[name]
      }
    }
  }
  return "no value"
}

//读取配置文件所有内容
func ReadList(c *Config) []map[string]map[string]string {
  file, err := os.Open(c.filepath)
  if err != nil {
    CheckErr(err)
  }
  defer file.Close()
  var data map[string]map[string]string
  var section string
  buf := bufio.NewReader(file)
  for {
    l, err := buf.ReadString('\n')
    line := strings.TrimSpace(l)
    if err != nil {
      if err != io.EOF {
        CheckErr(err)
      }
      if len(line) == 0 {
        break
      }
    }
    switch {
    case len(line) == 0:
    case line[0] == '[' && line[len(line)-1] == ']':
      section = strings.TrimSpace(line[1 : len(line)-1])
      data = make(map[string]map[string]string)
      data[section] = make(map[string]string)
    default:
      i := strings.IndexAny(line, "=")
      value := strings.TrimSpace(line[i+1 : len(line)])
      data[section][strings.TrimSpace(line[0:i])] = value
      if uniquappend(c, section) == true {
        c.conflist = append(c.conflist, data)
      }
    }

  }

  return c.conflist
}

func CheckErr(err error) string {
  if err != nil {
    return fmt.Sprintf("Error is :'%s'", err.Error())
  }
  return "Notfound this error"
}
//查找配置文件中是否存在所需的topic
func uniquappend(c *Config, conf string) bool {
	for _, v := range c.conflist {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}
/*func (c *Config) setValue(section, key, value string) bool {
  c.ReadList()
  date := c.conflist
  var ok bool
  var index = make(map[int]bool)
  var conf = make(map[string]map[string]string)
  for i, v := range data{
    _. ok = v[section]
    index[i] = ok
  }

  i, ok := func(m map[int]bool) (i int, v bool) {
    for i, v := range m {
      if v == true {
        return i, true
      }
    }
    return 0, false
  }(index)


    if ok {
    c.conflist[i][section][key] = value
    return true
  } else {
    conf[section] = make(map[string]string)
    conf[section][key] = value
    c.conflist = append(c.conflist, conf)
    return true
  }

  return false
}*/

