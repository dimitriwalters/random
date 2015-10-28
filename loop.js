/* The infamous loop problem in JavaScript */

function badLoop() {
    var counter = 0;
    console.log("Bad loop:");
    for (var i=0; i<10; i++) {
        setTimeout((function() {
            console.log(i);

            counter++;
            if (counter == 10) {
                goodLoop();
            }
        }));
    }
}

function goodLoop() {
    console.log("Good loop:");
    for (var i=0; i<10; i++) {
        setTimeout((function(i) {
            console.log(i);
        })(i));
    }
}

badLoop();