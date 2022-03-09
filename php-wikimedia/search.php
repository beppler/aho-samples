<?php

require __DIR__ . '/vendor/autoload.php';

use AhoCorasick\MultiStringMatcher;

$words = [];

$time_build = -microtime(true);

$handle = fopen("words.txt", "r");
while (($line = fgets($handle)) !== false) {
    $words[] = mb_strtolower(trim($line));
}
fclose($handle);

$matcher = new MultiStringMatcher($words);

$time_build += microtime(true);

$words_count = count($words);

$message = "I aguaranteeee the return to capital of this business cycle yields a lot of money related to bank loan swap with very basic risk on this base year.";

$time_match = -microtime(true);

$results = $matcher->searchIn(mb_strtolower($message));

$time_match += microtime(true);

echo "dictionary: count = {$words_count}\n";
echo "message: text = {$message}\n\n";


foreach ($results as $result) {
    echo "{$result[0]} -> {$result[1]}\n";
}

$time_build = round($time_build, 5);
$time_match = round($time_match, 5);
echo "\ntime: build = {$time_build}s, execute = {$time_match}s\n";

