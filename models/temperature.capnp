using Go = import "/go.capnp";
@0xc54e847d37419a66;
$Go.package("models");
$Go.import("github.com/alittlebrighter/thermometer-service/models");

enum TemperatureUnits {
    celsius @0;
    fahrenheit @1;
}

struct TemperatureRead {
    units @0 :TemperatureUnits;
    location @1 :Text;
    reading :union {
        value @2 :Float64;
        error @3 :Text;
    }
}