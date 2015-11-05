printList :: (Show a) => [a] -> IO()
printList [] = return ()
printList (x:xs) = do { print x ; printList xs }

sum' :: (Num a) => [a] -> a
sum' [] = 0
sum' (x:xs) = x + sum' xs

dot :: (Num a) => [a] -> [a] -> a
dot [] _ = 0
dot _ [] = 0
dot (x:xs) (y:ys) = x * y + dot xs ys