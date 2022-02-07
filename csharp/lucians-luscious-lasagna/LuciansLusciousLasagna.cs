class Lasagna
{
    // TODO: define the 'ExpectedMinutesInOven()' method
    public int ExpectedMinutesInOven(){
        return 40;
    }
    // TODO: define the 'RemainingMinutesInOven()' method
    public int RemainingMinutesInOven(int minutes_elapsed){
        return ExpectedMinutesInOven() - minutes_elapsed;
    }
    // TODO: define the 'PreparationTimeInMinutes()' method
    public int PreparationTimeInMinutes(int layer_count){
        return layer_count * 2;
    }
    // TODO: define the 'ElapsedTimeInMinutes()' method
    public int ElapsedTimeInMinutes(int layer_count, int minutes_elapsed){
        return PreparationTimeInMinutes(layer_count) + minutes_elapsed;
    }
}
