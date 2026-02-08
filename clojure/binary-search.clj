(ns binary-search)

(def numbers [1 3 5 6 9 10 12 22 23 24 26 29 30 31 32 37 38 39 40 42 44 45 46])

(defn binary-search
    "Returns the index of target in a sorted vector, or -1 if not found"
    [sorted target]
    (loop [left 0
        right (dec (count sorted))]
        (if (> left right)
            -1
            (let [mid (quot (+ left right) 2)
                mid-value (nth sorted mid)]
                (cond
                    (= mid-value target) mid
                    (< mid-value target) (recur (inc mid) right)
                    :else (recur left (dec mid)))))))

(println (binary-search numbers 22))
(println (binary-search numbers 19))
