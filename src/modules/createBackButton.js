const btnTmpl = require('../templates/button.pug')

const backButton = {
        name: 'Назад в меню',
        link: "main",
        classes: [
            'back_link',
        ],
    };

export default function createBackButton() {
    const button = btnTmpl(backButton);
    return button;
}