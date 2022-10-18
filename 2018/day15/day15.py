# from get_data import get_data_lines
# https://www.reddit.com/r/adventofcode/comments/a6chwa/2018_day_15_solutions/


MAX_HP = 200
ATTK = 3

SORT_MULTIPLIER = 1000


def unit_reading_sort_key(unit):
    return unit.y * SORT_MULTIPLIER + unit.x


class Unit:
    ELF = 'E'
    GOBLIN = 'G'

    def __init__(self, id, tp, hp, atk, x, y, unit_loc_map):
        self.id = id
        self.tp = tp
        self.hp = hp
        self.atk = atk
        self.x = x
        self.y = y
        self.alive = True
        unit_loc_map[(self.x, self.y)] = self

    def within_range(self, x, y):
        return (
            (self.y == y and (self.x == x + 1 or self.x == x - 1)) or
            (self.x == x and (self.y == y + 1 or self.y == y - 1))
        )

    def adjacent_to_enemy(self, units):
        adjacent = []
        for unit in units:
            if not unit.alive:
                continue
            if self.id == unit.id or self.tp == unit.tp:
                continue
            if self.within_range(unit.x, unit.y):
                adjacent.append(unit)
        # Sort by lowest HP, then reading order.
        return sorted(adjacent, key=lambda unit: (
                        unit.hp * SORT_MULTIPLIER * SORT_MULTIPLIER +
                        unit.y * SORT_MULTIPLIER + unit.x))

    def attack(self, unit, unit_loc_map):
        unit.hp -= self.atk
        if unit.hp <= 0:
            unit.die(unit_loc_map)

    def die(self, unit_loc_map):
        print(f'Unit {self} has died :(')
        del unit_loc_map[(self.x, self.y)]
        self.alive = False

    def find_targets_in_range(self, parsed_map, units):
        in_range = []
        for unit in units:
            if not unit.alive:
                continue
            if self.id == unit.id or self.tp == unit.tp:
                continue
            # consider the 4 pts in the cardinal directions from this unit.
            pts = check_in_range_pts(parsed_map, unit, units)
            in_range.extend(pts)
        return in_range

    def find_reachable(self, parsed_map, unit_loc_map, in_range):
        """
        this function finds the (x, y) points in in_range that are actually
        reachable from (self.x, self.y)
        """
        reachable = []
        for pt in in_range:
            r, cost, queue = self.is_reachable(parsed_map, unit_loc_map, pt)
            if r:
                reachable.append((pt, cost, queue))
        return reachable

    def other_unit_at(self, unit_loc_map, x, y):
        if (x, y) in unit_loc_map:
            unit = unit_loc_map[(x, y)]
            return unit.alive and unit.id != self.id

        return False

    def is_reachable(self, parsed_map, unit_loc_map, pt):
        # attempt to walk to pt. some sort of bucket fill algorithm?
        queue = [(pt[0], pt[1], 0)]
        queue_map = {(pt[0], pt[1]): 0}
        el_idx = -1
        while el_idx < len(queue) - 1:
            el_idx += 1
            x, y, ctr = queue[el_idx]
            # four closest cells
            ll = [(x + 1, y, ctr + 1),
                (x - 1, y, ctr + 1),
                (x, y + 1, ctr + 1),
                (x, y - 1, ctr + 1)]
            for cell in ll:
                x, y, cctr = cell[0], cell[1], cell[2]
                remove = False
                if parsed_map[y][x] == '#' or self.other_unit_at(
                        unit_loc_map, x, y):
                    remove = True
                if (x, y) in queue_map and queue_map[(x, y)] <= cctr:
                    remove = True
                if not remove:
                    queue.append(cell)
                    queue_map[(cell[0], cell[1])] = cell[2]

        if (self.x, self.y) in queue_map:
            return True, queue_map[(self.x, self.y)], queue

        return False, None, None

    def find_closest(self, reachable):
        reachable = sorted(
            reachable,
            key=lambda r: (r[1] * SORT_MULTIPLIER * SORT_MULTIPLIER +
                        r[0][1] * SORT_MULTIPLIER + r[0][0]))
        # sorted by cost, then Y, then X (reading order)
        # print(f'Closest sorted: {reachable}')
        return reachable[0]

    def move_to_closest(self, closest, unit_loc_map):
        # closest is a tuple. the second index is the queue.
        queue = closest[2]
        # find ourselves in the queue.
        our_cost = None
        for el in queue:
            if el[0] == self.x and el[1] == self.y:
                our_cost = el[2]
        # find the next item down costwise
        next_step = []
        for el in queue:
            if el[2] == our_cost - 1 and self.within_range(el[0], el[1]):
                next_step.append(el)

        next_step = sorted(
            next_step,
            key=lambda s: s[1] * 1000 + s[0])
        # print(f'Found next step: {next_step}')
        step = next_step[0]
        for s in next_step:
            assert abs(self.x - s[0]) + abs(self.y - s[1]) == 1

        del unit_loc_map[(self.x, self.y)]
        self.x = step[0]
        self.y = step[1]
        unit_loc_map[(self.x, self.y)] = self


def check_in_range_pts(parsed_map, unit, units):
    in_range = []
    for pt in ((unit.x+1, unit.y), (unit.x-1, unit.y),
            (unit.x, unit.y+1), (unit.x, unit.y-1)):
        if parsed_map[pt[1]][pt[0]] == '.':
            # reachable but check that there are no units here either.
            units_exist = False
            for u in units:
                if u.alive and u.x == pt[0] and u.y == pt[1]:
                    units_exist = True
                    break
            if not units_exist:
                in_range.append(pt)

    return in_range

def parse_map(battle_map, elf_attack_power):
    parsed_map = []
    units = []
    unit_loc_map = {}
    unit_id = 1
    for y, line in enumerate(battle_map):
        cur_line = []
        for x, chr in enumerate(line):
            if chr == '.' or chr == '#':
                cur_line.append(chr)
            elif chr == 'E' or chr == 'G':
                cur_line.append('.')
                unit = Unit(unit_id, chr, MAX_HP,
                            elf_attack_power if chr == 'E' else ATTK, x, y,
                            unit_loc_map)
                units.append(unit)
                print(f'Added unit {unit}')
                unit_id += 1

        parsed_map.append(cur_line)
    return parsed_map, units, unit_loc_map


def new_round(parsed_map, units, unit_loc_map):
    srt = sorted(units, key=unit_reading_sort_key)

    for unit in filter(lambda unit: unit.alive, srt):
        if (all_units_of_type_dead(units, Unit.ELF) or
                all_units_of_type_dead(units, Unit.GOBLIN)):
            return False  # round did not complete in its entirety.

        if not unit.alive:
            # The unit may have died mid-round.
            continue
        # 1) move AND attack!
        adjacent = unit.adjacent_to_enemy(units)
        if not adjacent:
            # move
            # find targets that are in range
            in_range = unit.find_targets_in_range(parsed_map, units)

            # print(f'In range of unit {unit}: ', in_range)
            reachable = unit.find_reachable(parsed_map, unit_loc_map, in_range)

            # print(f'Reachable from unit {unit}: ', reachable)
            if reachable:
                closest = unit.find_closest(reachable)

                # print(f'Chose a point: {closest}, moving to it')
                # finally, move to closest.
                if closest:
                    unit.move_to_closest(closest, unit_loc_map)

                    # we attack in the same turn!
                    adjacent = unit.adjacent_to_enemy(units)

        if adjacent:
            unit.attack(adjacent[0], unit_loc_map)

    return True


def all_units_of_type_dead(units, tp):
    return len(list(filter(lambda u: u.tp == tp and u.alive, units))) == 0


def setup_game(attack_power):
    battle_map = list(filter(lambda y: y.strip() != '', """
################################
###############.##...###########
##############..#...G.#..#######
##############.............#####
###############....G....G......#
##########..........#..........#
##########................##..##
######...##..G...G.......####..#
####..G..#G...............####.#
#######......G....G.....G#####E#
#######.................E.######
########..G...............######
######....G...#####E...G....####
######..G..G.#######........####
###.........#########.......E.##
###..#..#...#########...E.....##
######......#########.......####
#####...G...#########.....######
#####G......#########.....######
#...#G..G....#######......######
###...##......#####.......######
####..##..G........E...E..######
#####.####.....######...########
###########..#...####...E.######
###############...####..#...####
###############...###...#.E.####
#####################.#E....####
#####################.#...######
###################...##.#######
##################..############
##################...###########
################################
    """.split('\n')))

    # battle_map = get_data_lines(15)

    parsed_map, units, unit_loc_map = parse_map(battle_map, attack_power)
    print(f'Attack: {attack_power}')
    round_counter = 0
    while True:
        completed = new_round(parsed_map, units, unit_loc_map)
        if completed:
            round_counter += 1
            print(f'After {round_counter} rounds')
        else:
            print(f'Quitting after {round_counter} total rounds')
            break

    hp_sum = sum([u.hp for u in units if u.alive])
    print(f'Outcome: {round_counter * hp_sum}')
    return units


if __name__ == '__main__':
    # part1
    setup_game(3)

    # part2
    # attack_power = 2
    # while True:
    #     attack_power += 1
    #     units = setup_game(attack_power)
    #     success = True
    #     for unit in units:
    #         if unit.tp == Unit.ELF and unit.alive is False:
    #             success = False
    #     if success:
    #         print(f'Did it with attack_power {attack_power}')
    #         break