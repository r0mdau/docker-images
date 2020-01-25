#!/usr/bin/php
<?php

$email = "";
$stdin = fopen("php://stdin", "r");

while (!feof($stdin)) {
    $email .= fread($stdin, 1024);
}

fclose($stdin);
$filename = "Message_" . microtime(true) . "_" . mt_rand() . ".eml";
$file = fopen("/tmp/mail/$filename", "w");
fwrite($file, $email);
fclose($file);