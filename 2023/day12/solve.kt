https://todd.ginsberg.com/post/advent-of-code/2023/day12/

private fun countArrangements(
    springs: String,
    damage: List<Int>,
    cache: MutableMap<Pair<String, List<Int>>, Long> = HashMap()
): Long {
    val key = springs to damage
    cache[key]?.let {
        return it
    }
    if (springs.isEmpty()) return if (damage.isEmpty()) 1 else 0

    return when (springs.first()) {
        '.' -> countArrangements(springs.dropWhile { it == '.' }, damage, cache)

        '?' -> countArrangements(springs.substring(1), damage, cache) +
                countArrangements("#${springs.substring(1)}", damage, cache)


        '#' -> when {
            damage.isEmpty() -> 0
            else -> {
                val thisDamage = damage.first()
                val remainingDamage = damage.drop(1)
                if (thisDamage <= springs.length && springs.take(thisDamage).none { it == '.' }) {
                    when {
                        thisDamage == springs.length -> if (remainingDamage.isEmpty()) 1 else 0
                        springs[thisDamage] == '#' -> 0
                        else -> countArrangements(springs.drop(thisDamage + 1), remainingDamage, cache)
                    }
                } else 0
            }
        }

        else -> throw IllegalStateException("Invalid springs: $springs")
    }.apply {
        cache[key] = this
    }
}