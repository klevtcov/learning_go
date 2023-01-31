## Массивы и срезы в Go

_Массив_ в Go представляет собой структуру данных, состоящую из упорядоченной последовательности элементов, емкость которой определяется в момент создания. После определения размера массива его нельзя изменить. 
_Срез_ — это версия массива с переменной длиной, дающая разработчикам дополнительную гибкость использования этих структур данных. Срезы — это то, что обычно называют массивами в других языках.

### Массивы
Массивы представляют собой структурированные наборы данных с заданным количеством элементов. 
Массивы определяются посредством декларирования размера массива в квадратных скобках [ ], после которых указывается тип данных элементов. Все элементы массива в Go должны относиться к одному и тому же типу данных. После типа данных вы можете декларировать отдельные значения элементов массива в фигурных скобках { }.
> [capacity]data_type{element_values}

Если вы не декларируете значения элементов массива, по умолчанию используются нулевые значения, т. е. по умолчанию элементы массива будут пустыми. Это означает, что целочисленные элементы будут иметь значение 0, а строки будут пустыми.
> var numbers [3]int // [0 0 0]

```
coral := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
fmt.Println(coral)                           // [blue coral staghorn coral pillar coral elkhorn coral]
fmt.Printf("%q\n", coral)                    // ["blue coral" "staghorn coral" "pillar coral" "elkhorn coral"]
fmt.Println(coral[1])                        // staghorn coral
fmt.Println("Sammy loves " + coral[0])       // Sammy loves blue coral

coral[1] = "foliose coral"
fmt.Printf("%q\n", coral)                    // ["blue coral" "foliose coral" "pillar coral" "elkhorn coral"]
```
В Go имеется встроенная функция len(), помогающая работать с массивами и срезами. Как и в случае со строками, вы можете рассчитать длину массива или среза, используя команду len() с указанием массива или среза в качестве параметра.
```
len(coral) // 4
```
append() — это встроенный метод Go для добавления элементов в тип данных коллекции. _<b>Но данный метод не будет работать с помощью массива.</b>_ Как уже отмечалось, основное отличие массивов от срезов заключается в том, что размер массива нельзя изменить. Это означает, что хотя вы можете изменять значения элементов в массиве, вы не можете сделать массив больше или меньше после его определения.

