<?php
    function isLeap($yr) : bool {

        return $yr % 100 == 0 ? $yr % 400 == 0 : $yr % 4 == 0;

    }