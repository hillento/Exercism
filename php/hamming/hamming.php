<?php

function distance(string $strandA, string $strandB) : int
{
    $dif = 0;
   $strandA_len = strlen($strandA);
   $strandB_len = strlen($strandB);
    if ($strandA_len != $strandB_len)
    {
        throw new InvalidArgumentException('DNA strands must be of equal length.');
    }

    for ($i = 0; $i < $strandA_len; $i++)
    {
        if ($strandA[$i] != $strandB[$i]){
            $dif++;
        }
    }

    return $dif;
}