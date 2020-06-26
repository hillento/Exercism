using System;
using System.Collections.Generic;
using System.Linq;

public static class Pangram
{
    public static bool IsPangram(string input)
    {
        var alpha = "abcdefghijklmnopqrstuvwxyz".ToCharArray();
        //Turns original input to lower case
        input = input.ToLower();
        return alpha.All(input.Contains);
        
    }
}
