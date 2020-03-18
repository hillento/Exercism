<?php
function isPangram(string $input) {
 
    //gives an array containing all letters a-z
    $letters = range('a', 'z');
    //splits each input character into an array index for
    $input = str_split(strtolower($input));
    
    //If the letters array and input array contain 26 intersections (input has a value for each letter) then return true 
    
    return count(array_intersect($letters,$input)) == 26 ? true:false;
    
}