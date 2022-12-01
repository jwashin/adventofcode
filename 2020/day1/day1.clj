
(ns day1)

(slurp "input.txt")

(defn part1 [coll]
(for [x coll y coll
      :when (= (+ x y) 2020)
      :while (< y x)]
  (* x y)))

(defn part1-test []
  (let [coll [1721 979 366 299 675 1456]]
  (part1 coll)  )
  )
(println (part1-test))
