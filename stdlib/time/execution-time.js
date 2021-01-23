let start = Date.now();
for (let i = 0; i < 10000000000; i++) {}

let duration = (Date.now() - start) * 0.001;

console.log(duration, "s");
