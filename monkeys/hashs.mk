//hash测试示例
let people =[{"name":"Alice","age":24},{"name":"Anna","age":28}];

let getName = fn(person){person["name"];};

puts("-->");

puts(people[0]["name"]);

puts(people[1]["age"]);

puts(getName(people[1]));