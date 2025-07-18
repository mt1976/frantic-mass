package views

func selectUser() string {
	return `
		<div class="container">
			<h1>Select User</h1>
			<p>Please select a user to continue.</p>
			<form action="/choose_user" method="post">
				<select name="user" required>
					<option value="">-- Select User --</option>
					{{range .Users}}
						<option value="{{.ID}}">{{.Name}}</option>
					{{end}}
				</select>
				<button type="submit">Continue</button>
			</form>
		</div>
	`
}
