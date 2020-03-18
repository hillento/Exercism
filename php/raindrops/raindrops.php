<?php

function raindrops(int $num) : string {
    $sound = '';
    if($num % 3 == 0){
        $sound .= 'Pling';
    }
    if($num % 5 == 0){
        $sound .= 'Plang';
    }
    if($num % 7 == 0){
        $sound .= 'Plong';
    }
    return $sound != '' ? $sound : (string) $num;
}