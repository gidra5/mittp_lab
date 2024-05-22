# Результати досліджень

Перевірено декілька варіантів симуляції з використанням потоків на прикладі go.

1. Без будь якої сихронізації
2. Синхронізація на рівні кожної зміни
3. ... на рівні усього списку клітин
4. ... на рівні кожної з клітин що змінюються
5. ... на рівні лише одної з клітин що змінюються
6. Використовуючи паттерн producer/consumer

Результати симуляцій:
1.
```
Current cells state: 5 [5, 0, 0, 0, 0, 0, 0, 0, 0, 0]
Current cells state: 2278503 [66634, 269937, 207549, 240919, 270283, 379308, 404751, 192576, 221629, 25002]
Current cells state: 4721121 [132549, 546002, 418559, 510278, 571608, 782528, 830916, 408048, 468887, 51746]
Total cell count after simulation: 4721121
```

2. 
```
Current cells state: 5 [5, 0, 0, 0, 0, 0, 0, 0, 0, 0]
Current cells state: 1 [0, 0, 1, 0, 0, 0, 0, 0, 0, 0]
Current cells state: 3 [1, 0, 0, 0, 0, 0, 0, 2, 0, 0]
Current cells state: 3 [0, 1, 0, 0, 0, 1, 1, 0, 2, 0]
Total cell count after simulation: 5
```

3.
```
Current cells state: 5 [5, 0, 0, 0, 0, 0, 0, 0, 0, 0]
Current cells state: 5 [2, 0, 0, 1, 1, 0, 0, 1, 0, 0]
Current cells state: 5 [1, 1, 0, 0, 0, 1, 0, 1, 1, 0]
Total cell count after simulation: 5
```

4.
```
Current cells state: 5 [5, 0, 0, 0, 0, 0, 0, 0, 0, 0]
Current cells state: 5 [0, 0, 1, 1, 1, 1, 1, 0, 0, 0]
Total cell count after simulation: 5
```

5.
```
Current cells state: 5 [5, 0, 0, 0, 0, 0, 0, 0, 0, 0]
Current cells state: 378667 [104, 275971, 26406, 23, -294, 38176, 38659, -103, -357, 82]
Current cells state: 769728 [164, 556429, 54685, 6, -452, 80003, 79592, -229, -588, 118]
Total cell count after simulation: 769728
```

6.
```
Current cells state: 5 [5, 0, 0, 0, 0, 0, 0, 0, 0, 0]
Current cells state: 5 [0, 1, 2, 0, 0, 1, 0, 0, 0, 0]
Current cells state: 5 [0, 2, 0, 1, 0, 0, 1, 1, 0, 0]
Total cell count after simulation: 5
```