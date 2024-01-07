# protoc-gen-verifier

该仓库演示了如何进行protoc插件的开发。

protoc-gen-verifier可以根据指定的校验规则对message中的字段进行校验。
获取到每个字段上面的注释，所以为了简单起见，规定校验规则需要在字段的leading comments中指定。

支持的校验标签（tag）如下表所示：

<table align="left">
  <tr>
  </tr>     
  <tr>         
    <th style="border-right: 1px solid;border-left: 1px solid;"  >proto类型</th>
    <th style="border-right: 1px solid;border-left: 1px solid;" colspan="2"  > 支持的校验规则（tag）</th>
    <th style="border-right: 1px solid;border-left: 1px solid;" > 示例 </th>
  </tr>
  <tr>
    <td style="border-right: 1px solid;border-left: 1px solid;" rowspan="7"  > <b>int32 <br> int64 <br>  uint32 <br>  uint64 <br>  sint32 <br>  sint64 <br> fixed32 <br> fixed64 <br> sfixed32 <br> sfixed64</b> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;" > eq </td>
    <td style="border-right: 1px solid;"> 数值等于 </td>
    <td style="border-right: 1px solid;">  </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> gt </td>
    <td style="border-right: 1px solid;"> 数值大于 </td>
    <td style="border-right: 1px solid;"> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> gte </td>
    <td style="border-right: 1px solid;"> 数值大于等于 </td>
    <td style="border-right: 1px solid;">  </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> lt </td>
    <td style="border-right: 1px solid;"> 数值小于 </td>
    <td style="border-right: 1px solid;"> </td>
  </tr>

  <tr>
    <td style="border-right: 1px solid;"> lte </td>
    <td style="border-right: 1px solid;"> 数值小于等于 </td>
    <td style="border-right: 1px solid;">  </td>
  </tr>

  <tr>
      <td style="border-right: 1px solid;"> ne </td>
      <td style="border-right: 1px solid;"> 数值不等于 </td>
      <td style="border-right: 1px solid;">  </td>
    </tr>

  <tr>
    <td style="border-right: 1px solid;border-left: 1px solid;" rowspan="9"  > <b>string</b> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> len </td>
    <td style="border-right: 1px solid;"> 字符串长度等于 </td>
    <td style="border-right: 1px solid;"> len=11 </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> alpha </td>
    <td style="border-right: 1px solid;"> 字符串全部为字母 </td>
    <td style="border-right: 1px solid;">  </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> number </td>
    <td style="border-right: 1px solid;"> 字符串表示的内容为数字 </td>
    <td style="border-right: 1px solid;">  </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> contains </td>
    <td style="border-right: 1px solid;"> 字符串中包含子字符串 </td>
    <td style="border-right: 1px solid;"> contains=hello|world </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> startswith </td>
    <td style="border-right: 1px solid;"> 字符串以特定内容开头 </td>
    <td style="border-right: 1px solid;">  </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> endswith </td>
    <td style="border-right: 1px solid;"> 字符串以特定内容结尾 </td>
    <td style="border-right: 1px solid;">  </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> eq </td>
    <td style="border-right: 1px solid;"> 字符串相等 </td>
    <td style="border-right: 1px solid;"> eq=helloworld </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> ne </td>
    <td style="border-right: 1px solid;"> 字符串不相等 </td>
    <td style="border-right: 1px solid;">  </td>
  </tr>

  <tr>
    <td style="border-right: 1px solid;border-left: 1px solid;" rowspan="3"  > <b>bytes</b> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> len </td>
    <td style="border-right: 1px solid;"> 字节切片长度等于 </td>
    <td style="border-right: 1px solid;"> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> ne </td>
    <td style="border-right: 1px solid;"> 字节切片长度不等于 </td>
    <td style="border-right: 1px solid;"> </td>
  </tr>

  <tr>
    <td style="border-right: 1px solid;border-left: 1px solid;" rowspan="3"  > <b>repeated修饰</b> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> len </td>
    <td style="border-right: 1px solid;"> 切片长度等于 </td>
    <td style="border-right: 1px solid;"> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> ne </td>
    <td style="border-right: 1px solid;"> 切片长度不等于 </td>
    <td style="border-right: 1px solid;"> </td>
  </tr>

  <tr>
    <td style="border-right: 1px solid;border-left: 1px solid;" rowspan="3"  > <b>map</b> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> len </td>
    <td style="border-right: 1px solid;"> map中键值对数量等于 </td>
    <td style="border-right: 1px solid;"> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;"> ne </td>
    <td style="border-right: 1px solid;"> map中键值对数量不等于 </td>
    <td style="border-right: 1px solid;"> ne=6 </td>
  </tr>

  <tr>
    <td style="border-right: 1px solid;border-left: 1px solid;border-bottom: 1px solid;" rowspan="2"  > <b>bool</b> </td>
  </tr>
  <tr>
    <td style="border-right: 1px solid;border-bottom: 1px solid;"> eq </td>
    <td style="border-right: 1px solid;border-bottom: 1px solid;"> 布尔值相等 </td>
    <td style="border-right: 1px solid;border-bottom: 1px solid;"> eq=true或eq=false </td>
  </tr>
</table>

