<?php

/*
This is only a SKELETON file for the "Hamming" exercise. It's been provided as a
convenience to get you started writing code faster.

Remove this comment before submitting your exercise.
*/

function distance(string $strandA, string $strandB) : int
{
    $dif = 0;
   
    if (strlen($strandA) != strlen($strandB))
    {
        throw new InvalidArgumentException('DNA strands must be of equal length.');
    }

    for ($i = 0; $i < strlen($strandA); $i++)
    {
        if ($strandA[$i] != $strandB[$i]){
            $dif++;
        }
    }

    return $dif;
}
