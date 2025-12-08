use rand::Rng;

struct Data {
    randoms: Vec<u8>,
    input: Vec<u8>,
}

fn gen_randoms(n: usize) -> Vec<u8> {
    let mut rng = rand::thread_rng();
    (0..n).map(|_| rng.gen_range(0..100)).collect()
}

fn count_matches(randoms: &[u8], input: &[u8]) -> (Vec<usize>, usize) {
    let counts: Vec<usize> = input
        .iter()
        .map(|&v| randoms.iter().filter(|&&x| x == v).count())
        .collect();
    let total: usize = counts.iter().sum();
    (counts, total)
}

fn main() {
    // 問題2: 関数 gen_randoms を使って配列を作成
    let mut data = Data {
        randoms: gen_randoms(10), // randoms に 0..99 の乱数を10個
        input: gen_randoms(3),    // input に 0..99 の乱数を3個
    };

    // 問題1: 最初の配列と一致数を表示
    println!("initial randoms: {:?}", data.randoms);
    println!("input: {:?}", data.input);
    let (counts, total) = count_matches(&data.randoms, &data.input);
    for (i, &v) in data.input.iter().enumerate() {
        println!("value {} appears {} time(s) in randoms", v, counts[i]);
    }
    println!("total matches: {}", total);

    // 問題3: randoms に input の数字が全部含まれるまで繰り返す
    let mut iterations = 0usize;
    loop {
        iterations += 1;
        let (counts, _) = count_matches(&data.randoms, &data.input);
        if counts.iter().all(|&c| c > 0) {
            // 全て含まれた
            break;
        }
        // 含まれていない場合は randoms を再生成して試行を繰り返す
        data.randoms = gen_randoms(10);
    }

    println!("\nfinal randoms: {:?}", data.randoms);
    println!("input: {:?}", data.input);
    let (final_counts, final_total) = count_matches(&data.randoms, &data.input);
    for (i, &v) in data.input.iter().enumerate() {
        println!("value {} appears {} time(s) in final randoms", v, final_counts[i]);
    }
    println!("final total matches: {}", final_total);
    println!("iterations until all input found: {}", iterations);
}

