<?php

function from(DateTimeImmutable $date) : DateTimeImmutable {

    $gigasec = 1e9;

    $date = $date->add(new DateInterval('PT'.$gigasec.'S'));
    return $date;


}