select FirstName, LastName, City, State from Person left join Address on Person.PersonId = Address.PersonId;

# 使用Using来声明使用哪个列名来进行联合
SELECT FirstName,LastName,City FROM Person LEFT JOIN Address USING (PersonId);

