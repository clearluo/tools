/*
 * 程序崩溃邮件告警脚本
 * 
 * crontab配置:  */1 * * * * /usr/bin/php -f /data/core/monitor.php  >> /data/core/monitor.log 2>&1
 */
<?php

$file_path = "/data/core/";

$cmd_str = "cat $file_path"."curr.data";
$out_max = shell_exec($cmd_str);
$curr_max = $out_max ? $out_max : 0;
echo "curr_max = ".$curr_max;

$cmd_str = "ls -rt $file_path"."core*";
exec($cmd_str, $out_file_name, $status);
print_r($out_file_name);

$time_name = array();

foreach($out_file_name as $key=>$value){
	$arg_name = explode("-", $value);
	$time_name[] = $arg_name[6];
}

print_r($time_name);

$flags = false;
foreach ($time_name as $key=>$value){
	if ($value > $curr_max){
		$content = date("Y-m-d H:i:s", $value) . " happen to core of " . $out_file_name[$key];
		$cmd_str = "echo $content |mail -s '程序崩溃'  tolsw@qq.com";
		echo $cmd_str . "\n";
		shell_exec($cmd_str);
		$curr_max = $value;
        $false = true;
	}
}

if ($flag){
    $cmd_str = "echo $curr_max > $file_path"."curr.data";
    echo "cmd_str = ". $cmd_str . "\n";
    shell_exec($cmd_str);
}



