<?php

function toRoman($num){
    
    $out = '';

    $possible = [
        1000 => 'M',
        900  => 'CM',
        500  => 'D',
        400  => 'CD',
        100  => 'C',
        90   => 'XC',
        50   => 'L',
        40   => 'XL',
        10   => 'X',
        9    => 'IX',
        5    => 'V',
        4    => 'IV',
        1    => 'I'
    ];

    foreach($possible as $digits=>$roman){
        while($num >= $digits){
            $num -= $digits;
            $out .= $roman;
        }
    }
    return $out;
}
