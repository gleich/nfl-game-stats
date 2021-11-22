import csv


def main():
    data = read_csv()
    categorized_data = categorize(data)
    output(categorized_data)


def read_csv():
    data = []
    with open("data.csv") as csv_file:
        csv_reader = csv.reader(csv_file)
        next(csv_reader)
        for line in csv_reader:
            if line[28] == "" or line[29] == "":
                continue
            data.append((int(line[1]), abs(int(line[28]) - int(line[29]))))
    return data


def categorize(data):
    # year > range
    categorized_data = {}
    # setting initial values
    for game in data:
        season = game[0]
        if categorized_data.keys().__contains__(game[0]):
            continue
        else:
            categorized_data[season] = {}
            for i in range(5, 55, 5):
                categorized_data[season][i] = 0
    # grouping scores
    for game in data:
        season = game[0]
        score = game[1]
        for i in range(5, 55, 5):
            if i == 50 and score > i:
                categorized_data[season][i] += 1
            elif score >= i - 5 and score <= i:
                categorized_data[season][i] += 1
    return categorized_data


def output(data):
    rows = [""]
    for range_value in data[1920].keys():
        rows[0] += f" & \leq {range_value}"
    for year in data:
        row = f"{year}"
        for range_value in data[year]:
            row += f" & {data[year][range_value]}"
        rows.append(row)
    print(" \\\\ \\hline\n".join(rows))


if __name__ == "__main__":
    main()
