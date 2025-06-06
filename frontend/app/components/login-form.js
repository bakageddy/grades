import Component from "@glimmer/component";
import { tracked } from "@glimmer/tracking";
import { action } from "@ember/object";
import { service } from "@ember/service";

export default class LoginForm extends Component {
  @tracked register_no = '';
  @service router;

  get register_no() {
    this.register_no;
  }

  @action
  handle_input_register(ev) {
    this.register_no = ev.target.value;
    console.log(this.register_no);
  }

  @action
  async handle_submit(_) {
    let params = new URLSearchParams({
      "reg_no": this.register_no,
      "password": this.password
    }).toString();

    let resp = await fetch("https://localhost:42069/auth/login?" + params, {
      "method": "POST"
    });

    if (resp.status !== 200) {
      alert("Something went wrong! Try again later ( P-P)");
    }

    this.router.transitionTo('app/home')
  }
}
