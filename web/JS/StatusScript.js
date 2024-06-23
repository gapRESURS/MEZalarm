const ipElements = Array.from(document.querySelectorAll('div[id="ping"]'));

async function pingIP(element) {
    const ip = element.getAttribute('data-ip');

    try {
        const response = await fetch(`http://192.168.14.192/SKUD/ping/?ip=${ip}`);
        const data = await response.json();

        if (data.result === true) {
            // IP адрес доступен, меняем цвет текста на зеленый (Bootstrap: success)
            element.classList.remove('text-danger');
            element.classList.add('text-success');
        } else {
            // IP адрес недоступен, меняем цвет текста на красный (Bootstrap: danger)
            element.classList.remove('text-success');
            element.classList.add('text-danger');
        }
    } catch (error) {
        console.error('Ошибка при выполнении пинга:', error);
    }
}

setInterval(() => {
    ipElements.forEach(element => pingIP(element));
}, 5000);