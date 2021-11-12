const submitBtn = document.getElementById('submitBtn');
const cityName = document.getElementById('cityName');
const city_name = document.getElementById('city_name');
const temp_status = document.getElementById('temp_status');
const temp_real_val = document.getElementById('temp_real_val');
const datahide = document.querySelector('.middle_layer');
const getInfo = async (event) => {
    event.preventDefault();
    let cityVal = cityName.value;

    if (cityVal == "") {
        city_name.innerText = `Pls write the word befor search`;
        datahide.classList.add('data_hide');
    } else {
        try {
            let url = `http://still-stream-79080.herokuapp.com/query?word=${cityVal}`;
            const response = await fetch(url);
            const data = await response.json();
            
            let answer = JSON.stringify(data);
            let pk = ""
            let k = 0
            for (let i = 0; i < answer.length; i++) {
                k++;
                if (answer[i] == '{' || answer[i] == '}' || answer[i] == '[' || answer[i] == ']' || answer[i] == '"') {
                    continue;
                } else {
                    pk += answer[i];
                }
                if (answer[i] == ',') answer += " ";
                if ((i%30 == 0) || (k > 25 && answer[i] == ',')) {
                    pk += "\n";
                    k = 0;
                }
            }
            let lk = pk.substring(6, pk.length)
            city_name.innerText = `${lk}`;
        }catch {
            //console.log(Error);
            city_name.innerText = `Pls Enter valid word name ${Error}`;
            datahide.classList.add('data_hide');
        }
    }
}
submitBtn.addEventListener('click', getInfo);