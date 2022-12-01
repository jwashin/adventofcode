(ns day1.core
  (:require [clojure.string :as str])
  (:gen-class))

(defn part1 [coll]
  (for [x coll y coll
        :when (= (+ x y) 2020)
        :while (< y x)]
    (* x y)))

(defn part2 [coll]
  (for [x coll y coll z coll
        :when (= (+ x y z) 2020)]
    (* x y z)))

(comment
  (defn part1-test []
    (let [coll [1721 979 366 299 675 1456]]
      (part1 coll)))
  (println (part1-test)))

(def inpt (map #(Integer/parseInt %1) (str/split
                                       (str (slurp "input.txt"))
                                       #"\n")))
;; inpt

(println "Part 1: " (first (part1 inpt)))
(println "Part 2: " (first (part2 inpt)))