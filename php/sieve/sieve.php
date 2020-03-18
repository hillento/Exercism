<?php
function sieve($num) {

    $range = $num>=2 ? range(2, $num): [];

    $range_l = count($range);

    for($j=0; $j<$range_l; $j++){
        for($i=1; $i<$range_l; $i++){
            if($range[$i] % $range[$j] == 0 && $range[$i] != $range[$j]){
                array_splice($range, $i,1);
                $range_l = count($range);
            };
        };
    };
    return $range;
}