import Component from "@glimmer/component";
import { tracked } from "@glimmer/tracking";
import { action } from "@ember/object";

export default class LoginForm extends Component {
  @tracked register_no;
  @tracked password;

  @action
  handle_input_register(ev) {
    this.register_no = ev.target.value;
    console.log(this.password);
  }

  @action
  handle_input_password(ev) {
    this.password = ev.target.value;
    console.log(this.password);
  }

  @action
  async handle_click(_) {
    let params = new URLSearchParams({
      "reg_no": this.register_no,
      "password": this.password
    }).toString();
    let resp = await fetch("https://localhost:42069/auth?" + params, {
      "method": "POST"
    });

    if (resp.status !== 200) {
      alert("Something went wrong! Try again later ( P-P)");
    }
  }
}
