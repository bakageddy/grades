import Component from "@glimmer/component";
import { action } from "@ember/object";
import { tracked } from "@glimmer/tracking";
import { service } from "@ember/service";

export default class RegisterFormComponent extends Component {
  @tracked register_no;
  @service router;

  get register_no() {
    return this.register_no;
  }

  @action
  handle_input_register(ev) {
    this.register_no = ev.target.value;
    console.log(this.register_no);
  }

  @action
  async handle_submit(_) {
    let params = new URLSearchParams({
      "register_no": this.register_no,
    }).toString();

    let resp = await fetch("https://localhost:42069/auth/register?" + params, {
      "method": "POST",
    });

    if (resp.status !== 200) {
      alert("Something went wrong! Try again later ( P-P)");
    }

    this.router.transitionTo('login')
  }
}
