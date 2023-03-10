## JSON

### Общие сведения о формате JSON

JSON (JavaScript Object Notation) – текстовый формат обмена структурированными данными, основанный на JavaScript. JSON – не единственная доступная нам форма решить данную задачу: аналогичной цели служат XML, YAML и Google’s Protocol Buffers, и каждый имеет свою нишу, однако из-за простоты, удобочитаемости и всеобщей поддержки наиболее широко используется именно JSON.

JSON представляет собой одну из двух структур:

+ набор пар ключ - значение;
+ упорядоченный набор значений.

В качестве значений в JSON могут быть использованы:

+ объект – это неупорядоченное множество пар ключ - значение, заключённое в фигурные скобки «{ }». Ключ описывается строкой, между ним и значением стоит символ «:». Пары ключ - значение отделяются друг от друга запятыми;
+ массив — это упорядоченное множество значений, заключенных в квадратные скобки «[ ]». Значения разделяются запятыми;
+ число (целое или вещественное);
+ литералы true, false (логические) и null;
+ строка.

Рассмотрим пример данных в формате JSON:
```
{
    "firstName": "Иван",
    "lastName": "Иванов",
    "address": {
        "streetAddress": "Московское ш., 101, кв.101",
        "city": "Ленинград",
        "postalCode": 101101
    },
    "phoneNumbers": [
        "812 123-1234",
        "916 123-4567"
    ]
}
```
Здесь данные представляют из себя набор пар ключ - значение, при этом: firstName и lastName - строки, address - вложенный набор пар ключ - значения, а phoneNumbers - упорядоченный набор значений.

Создание структуры для GO из JSON
+ https://mholt.github.io/json-to-go/

### Кодирование и декодирование данных

Marshal и Unmarshal (кодирование и декодирование) данных в формате JSON в стандартной библиотеке Go реализовано в пакете encoding/json.

Наиболее удобным типом для кодирования / декодирования таких данных является структура и срез структур, именно его мы и рассмотрим в рамках этого курса. Но при этом нужно отметить, что в некоторых случаях, когда структура данных нам не известна, мы можем декодировать данные в типы с использованием интерфейсов: interface{}, map[string]interface{}, []interface{}, []map[string]interface{}, однако в дальнейшем такой способ потребует использования рефлексии для анализа таких данных, а это выходит за пределы рассматриваемой темы.

В рассматриваемом пакете мы можем найти 3 функции, позволяющие кодировать / декодировать данные в байтовый срез, чтобы иметь возможность рассматривать конкретные примеры работы, давайте сначала рассмотрим эти функции. Начнем с функции для кодирования данных Marshal:
```
type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

s := myStruct{
	Name:   "John Connor",
	Age:    35,
	Status: true,
	Values: []int{15, 11, 37},
}

// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
// и возвращает байтовый срез с данными, кодированными в формат JSON.
data, err := json.Marshal(s)
if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("%s", data) // {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
```
Как мы видим, данные закодированы и мы даже можем их прочитать. Если же мы хотим получить результат, который лучше подходит именно для чтения человеком (например для использования в качестве конфигурационного файла или для отображения информации на экране компьютера, а не для передачи данных другой программе по сети), мы можем использовать родственную функцию MarshalIndent.

MarshalIndent похож на Marshal, но применяет отступ (indent) для форматирования вывода. Каждый элемент JSON в выходных данных начинается с новой строки, начинающейся с префикса (prefix), за которым следует один или несколько отступов в соответствии с вложенностью:
```
type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

s := myStruct{
	Name:   "John Connor",
	Age:    35,
	Status: true,
	Values: []int{15, 11, 37},
}

data, err := json.MarshalIndent(s, "", "\t")
if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("%s", data)

//{
//	"Name": "John Connor",
//	"Age": 35,
//	"Status": true,
//	"Values": [
//		15,
//		11,
//		37
//	]
//}
```
Ну и в завершении этого шага рассмотрим последнюю из трех функций Unmarshal, она принимает в качестве аргумента байтовый срез и указатель на объект, в который требуется декодировать данные. Рассмотрим это на уже знакомом примере:
```
data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

var s myStruct

if err := json.Unmarshal(data, &s); err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("%v", s) // {John Connor 35 true [15 11 37]}
```

### Проверка json на правильность

Мы можем проверить является ли срез байтов форматом json:
```
type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

m := user{Name: "John Connor", Email: "test email"}

data, _ := json.Marshal(m)

data = bytes.Trim(data, "{") // испортим json удалив '{'

// функция json.Valid возвращает bool, true - если json правильный
if !json.Valid(data) {
	fmt.Println("invalid json!") // вывод: invalid json!
}

fmt.Printf("%s", data) // вывод: "Name":"John Connor","Email":"test email","Status":false,"Language":null}
```

### Аннотирование структур
Давайте рассмотрим вопрос тонкой настройки кодирования / декодирования данных в формате JSON. На предшествующих шагах мы видели, что имена полей объектов JSON выглядят также, как и имена полей структуры, в которую или из которой они кодируются. Но что если мы хотим изменить эти имена? Или кодировать / декодировать только часть полей? Для этого мы можем использовать специальные аннотации - тэги для полей нашей структуры.

Продемонстрировать используемые теги проще всего на примере:
```
// в общем виде аннотация выглядит так: `json:"используемое_имя,опция,опция"`

type myStruct struct {
	// при кодировании / декодировании будет использовано имя name, а не Name
	Name string `json:"name"`
	
	// при кодировании / декодировании будет использовано то же имя (Age),
	// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
	// то при кодировании оно будет опущено
	Age int `json:",omitempty"`
	
	// при кодировании / декодировании поле всегда игнорируется
	Status bool `json:"-"`
}

m := myStruct{Name: "John Connor", Age: 0, Status: true}

data, err := json.Marshal(m)
if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("%s", data) // {"name":"John Connor"}
```
Как видите, в закодированных в формат JSON данных поле "Name" именуется как "name", а Age и Status отсутствуют, но по разным причинам: Status всегда игнорируется, поскольку установлен тег "-", а Age проигнорирован, т.к. его значение 0 и установлен тег "omitempty". Таким образом, мы можем довольно тонко настроить процесс кодирования / декодирования данных: использовать требуемые нам имена, игнорировать "пустые" (нулевые) значения для экономного использования ресурсов или же просто игнорировать определенные поля структур, которые в работе не будут использоваться.
 

Завершая рассмотрение этого вопроса нужно отметить следующее: неэкспортируемые поля (имена которых начинаются со строчной буквы) не участвуют в кодировании / декодировании.

### Типы Encoder и Decoder
encoding/json позволяет оперировать не только байтовыми срезами, но и работать с уже знакомыми нам типами io.Reader и io.Writer, для этого пакет предоставляет нам типы Encoder и Decoder. Данные типы помимо методов Encode() и Decode() предоставляют нам ряд дополнительных методов, которые могут быть использованы в определенных случаях. Рассмотрим базовые методы, которые применяются наиболее часто:
```
type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	src = testStruct{Name: "John Connor", Age: 35} // структура с данными
	dst = testStruct{}                             // структура без данных
	buf = new(bytes.Buffer)                        // буфер для чтения и записи
)

enc := json.NewEncoder(buf)
dec := json.NewDecoder(buf)

enc.Encode(src)
dec.Decode(&dst)

fmt.Print(dst) // {John Connor 35}
```
В примере мы сделали следующее: во-первых, создали объекты Encoder и Decoder с помощью функций NewEncoder и NewDecoder соответственно. Каждая из этих функций получила в качестве аргумента буфер, который удовлетворяет одновременно и интерфейсу io.Reader, и интерфейсу io.Writer, соответственно мы смогли сначала записать в него данные, а затем их прочитать, используя методы Encode и Decode соответственно.

В каких ситуациях применять функции Marshal и Unmarshal и методы Encode и Decode соответственно? Все зависит от того, с каким типом данных мы работаем, какой в настоящее время нам более удобен. Обратите внимание, в отличие от данных в формате CSV, с которыми мы работали ранее, данные в формате JSON являются цельным объектом (в т.ч. из-за используемых скобок), поэтому мы не можем кодировать и декодировать их поэтапно.
