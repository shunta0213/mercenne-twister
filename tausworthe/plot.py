import matplotlib.pyplot as plt

# テキストファイルからデータを読み込む
for i in range(25):
    file_path = f'./data/seed_random_{i}.txt'
    with open(file_path, 'r') as file:
        data = [float(line.strip()) for line in file]

    # ヒストグラムを作成
    plt.hist(data, bins=20, edgecolor='black')
    plt.title('Histogram of Random Numbers')
    plt.xlabel('Value')
    plt.ylabel('Frequency')
    plt.grid(True)
    plt.savefig(f"fig/seed_randomness_{i}.png")
    plt.close()
