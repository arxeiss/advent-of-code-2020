<?php

libxml_use_internal_errors(true);

$data = file_get_contents("task.html");

$dom = new DOMDocument();
$dom->loadHtml(mb_convert_encoding($data, 'HTML-ENTITIES', 'UTF-8'));
$finder = new DomXPath($dom);

$desc = $finder->query("//*[@class='day-desc']");

ob_start();
$header = $finder->query(".//h2", $desc->item(0));
$headerText = strip_tags($dom->saveHtml($header->item(0)));
preg_match('/Day ([0-9]+)/i', $headerText, $matches);

echo "# {$headerText}\n\n";
echo "See [How to run](https://github.com/arxeiss/advent-of-code-2020/#how-to-run) chapter to run this puzzle\n\n";
echo "> :warning: **SPOILER ALERT** :warning: - The code contains solution for the whole task. Try first to solve it **yourself**. :link: https://adventofcode.com/2020/day/{$matches[1]}\n\n";

echo "## --- Part 1 ---\n\n";

$content = $dom->saveHtml($desc->item(0)).$dom->saveHtml($desc->item(1));
$content = preg_replace_callback('/<pre[^>]*>\s*<code[^>]*>(.*?)<\/code>\s*<\/pre>\s*/is', function ($matches) {
  $code = strip_tags(trim($matches[1]));
  return "```\n{$code}\n```\n\n";
}, $content);
$content = preg_replace_callback('/<code[^>]*>(.*?)<\/code>/is', function ($matches) {
  return '`'.strip_tags(trim($matches[1])).'`';
}, $content);
$content = preg_replace('/<\/?em[^>]*>/i', '**', $content);
$content = preg_replace('/<li[^>]*>(.*?)(<\/li>|\n)\s*/i', "- $1\n", $content);
$content = preg_replace('/<\/(p|ul)>\s*/i', "\n\n", $content);
$content = preg_replace('/<h2[^>]*>(.*?)<\/h2>\s*/i', "## $1\n\n", $content);
$content = preg_replace('/[\n]{3,}/', "\n\n", strip_tags($content));

echo "{$content}\n";
file_put_contents("task.md", ob_get_clean());
