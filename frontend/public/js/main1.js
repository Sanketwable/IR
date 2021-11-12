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
        city_name.innerText = `Pls write the name befor search`;
        datahide.classList.add('data_hide');
    } else {
        try {
            let url = `https://still-stream-79080.herokuapp.com/addstr?word=${cityVal}`;
            const response = await fetch(url);
            let lk = 'Word Added'
            city_name.innerText = `${lk}`;
        }catch {
            //console.log(Error);
            city_name.innerText = `Word Added`;
            datahide.classList.add('data_hide');
        }
    }
}
submitBtn.addEventListener('click', getInfo);