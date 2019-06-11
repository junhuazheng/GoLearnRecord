xml
package
import "encoding/xml"

函数
Marshal
func Marshal(v interface{})([]byte, error)
Marshal函数返回v的XML编码。
Marshal处理数组或者切片时会序列化每一个元素。Marshal处理指针时，会序列化其指向的值；
如果指针为nil，则啥也不会输出。Marshal处理接口时，会序列化其内包含的具体类型值，如果
接口为nil，也是不会输出。Marshal处理其余类型数据时，会输出一个或多个包含数据的XML元素。

Unmarshal
func Unmarshal(data []byte, v interface{}) error
Unmarshal解析XMl编码的数据并将其结果存入v指向的值，V只能指向结构体、切片或者字符串。
良好格式化的数据如果不能存入v，会被丢弃。
以为Unmarshal使用reflect包，它只能填写导出字段。本函数好似用大小写敏感的比较来匹配
XML元素名和结构体的字段名/标签键名。


NewEncoder
func NewEncoder(w io.Writer) *Encoder
(*Encoder)Encode
func (enc *Encoder) Encode(v interface{}) error
Encode将v编码为XML后写入底层。

NewDecoder
func NewDecoder(r io.Reader) *Decoder
创建一个从r读取XMl数据的解析器。如果r未实现io.ByteReader接口，NewDecoder会为其添加缓存。
(*Decoder) Decoder
func (d *Decoder) Decode(v interface{}) error
Decode方法功能类似xml.Unmarshal函数，但是会从底层读取XML数据并查找StartElement.


Token
除了上面的方法，xml包还提供了其他的解析xml方法。
(*Decoder) Token
func (d *Doceder) Token() (t Token, err error)Token
返回输入流里的下一个XML token。在输入流的结尾处，会返回(nil, io.EOF)
返回的token数据里面的[]byte数据引用自解析器内部的缓存，只在下一次调用Token之前有效。
如果要获取切片的拷贝，调用CopyToken函数或者token的Copy方法。
成功调用的Token方法会将自我闭合的元素扩展为分离的起始和结束标签。
Token方法会保证它返回的StartElement和EndElement两种token正确的嵌套和匹配：
如果本方法遇到了不正确的结束标签，会返回一个错误。


json
package
import "encoding/json"

函数
Marshal
func Marshal(v interface{}) ([]byte, error)
Marshal函数返回v的json编码。
Marshal函数会递归的处理值。如果一个值实现了Marshaler接口且非nil指针，会调用其MarshalJSON方法来生成json编码。nil指针异常并不是严格必须的，但会模拟与UnmarshalJSON行为类似的必须和异常。否则，Marshal函数使用下面的基于类型的默认编码格式：
布尔类型编码为josn布尔类型
浮点数、整数和Number类型的值编码为josn数字类型
字符串编码为json字符串。角括号"<"和">"会转义为"\u003c"和"\u003e"以避免某些路蓝旗把json输出错误理解为HML。基于同样的原因，"&"转义为"\u0023"。
数组和切片类型的值编码为josn数组，但[]byte编码为base64编码字符串，nil切片编码为null。

Unmarshal
func Unmarshal(data []byte, v interface{}) error
Unmarshal函数解析json编码的数据并将结果存入v指向的值
Unmarshal和Marshal做相反的操作，必要时申请映射、切片或指针，有如下的附加规则：
要将json数据解码写入一个指针，Unmarshal函数首先处理json数据是json字面值null的情况。此时，函数将指针设为nil；否则，函数将json数据解码写入指针指向的值；如果指针本身是nil，函数会先申请一个值并使指针指向它。
要将json数据解码写入一个结构体，函数会匹配输入对象的键和Marshal使用的键（结构体字段名或者它的标签制定的键名），优先选择精确的匹配，但也接受大小写不敏感的匹配。

要将json数据解码吸入一个接口类型值，函数会将数据解码为如下类型写入接口:
Bool                    对应JSON布尔类型
float64                 对应JSON数字类型
string                  对应JSOn字符串类型
[]interface{}           对应JSON数组
map[string]interface{}  对应JSON对象
nil                     对应JSON的null
如果一个JSON值不匹配给出的目标类型，或者如果一个json数字写入目标类型时一出，Unmarshal函数会跳过该字段并尽量完成其余的解码操作。如果没有出现更加严重的错误，本函数会返回一个描述第一个此类错误的详细信息的UnmarshalTypeError。
JSON的null值解码为go的接口、指针、切片时会将他们设置为nil，因为null在json里面一般表示“不存在”。解码json的null值到其他go类型时，不会造成任何改变，也不会产生错误。
当解码字符串时，不合法的utf-8或者utf-16代理（字符）对不失为错误，而是将非法字符替换为unicode字符U+FFFD。