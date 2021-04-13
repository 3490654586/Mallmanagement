<?

$id = $_GET['x'];//接受get传递的参数名x的值并赋值给变量id
$conn = mysql_connect('localhost','root','');//连接mysql数据库
mysql_select_db('zuhao',$conn);//选择$conn连接请求下的fanke数据库名
$sql = "select * from user where id=$id"; //定义sql语句并组合变量id
$result = mysql_query($sql);//执行sql语句并返回给变量result
while($row = mysql_fetch_array($result)){ //遍历数组数据并显示
	echo "用户ID：".$row['id']."<br >";
	echo "用户名：".$row['username']."<br >";
	echo "用户密码：".$row['password']."<br >";
}
mysql_close($conn);//关闭数据库连接
echo "<hr>";
echo "你当前执行的sql语句为：";
echo "select * from user where id=$id";


?>