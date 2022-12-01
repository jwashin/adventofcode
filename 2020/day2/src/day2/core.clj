(ns day2.core
  (:require [clojure.string :as str])
  (:gen-class))

(comment
  (def input1 (str/split (str (slurp "input.txt"))
                         #"\n"))
  
  
  
  
  
  
  )
input1

(parse-long "123122312")


(defn- parsing-err
  "Construct message for parsing for non-string parsing error"
  ^String [val]
  (str "Expected string, got " (if (nil? val) "nil" (-> val class .getName))))

(defn parse-long
  {:doc "Parse string of decimal digits with optional leading -/+ and return a
  Long value, or nil if parse fails"
   :added "1.11"}
  ^Long [^String s]
  (if (string? s)
    (try
      (Long/valueOf s)
      (catch NumberFormatException _ nil))
    (throw (IllegalArgumentException. (parsing-err s)))))


(def input (str/split (str (slurp "input.txt"))
                      #"\n"))

(def test-data (str/split "1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc" #"\n"))

;; (def rgx #"(\d)-(\d)\s([a-z])\:\s([a-z].)")



(def s-parser #"(?<f>\d)-(?<t>\d)\s(?<c>[a-z]):\s(?<test>[a-z]+)")

s-parser
test-data
input



(def matcher (re-matcher s-parser (first test-data)))
(re-find matcher)

(.group matcher "f"                                                                                             )
(.group matcher "test")

(matcher (first test-data))
(re-find (matcher (first test-data)))
(first (re-find (matcher (first test-data))))
                                        
(defn is= [a b] 
  (= a b)
  )

(is= 3 3)



(let [patt (re-pattern "(?<area>\\d{3})-(?<prefix>\\d{3})-(?<tail>\\d{4})")]
  ; `re-matches` will find the capturing groups and stick them in a vector
  ; after the full match The capture groups are numbered starting with 1.
  ; The full match is like group zero.
  (is= ["619-239-5464" "619" "239" "5464"] (re-matches patt "619-239-5464"))

  ; Construct a java.util.regex.Matcher.  Keep in mind that it is a mutable object!
  (let [matcher (re-matcher patt "619-239-5464")]
    ; Execute the Matcher via `re-find`. It returns all 4 groups and caches them
    (is= ["619-239-5464" "619" "239" "5464"] (re-find matcher))

    ; `re-groups` simply returns the cached result from the Matcher
    (is= ["619-239-5464" "619" "239" "5464"] (re-groups matcher))

    ; We need the instance function Matcher.group( <name> ) to
    ; extract the named group
    (is= "619" (.group matcher "area"))
    (is= "239" (.group matcher "prefix"))
    (is= "5464" (.group matcher "tail"))))





(defn valid1 [s])


(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (println "Hello, World!"))


