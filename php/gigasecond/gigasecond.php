<?php

function from(DateTimeImmutable $date) {
    //Seconds in a gigasecond
    $gs = 1000000000;
    return $date->add(new DateInterval('PT'.strval($gs).'S'));


}