import PropTypes from 'prop-types';
import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { authenticationService } from './services/authentication';

const styles = theme => ({
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200,
  },
  button: {
    margin: theme.spacing.unit,
  },
});

class Login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      name: '',
      password: '',
    };

    // redirect to home if already logged in
    if (authenticationService.currentUserValue) {
      this.props.history.push('/');
    }

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange (event) {
    switch (event.target.name) {
      case 'name':
        this.setState({name: event.target.value});
        break;
      case 'password':
        this.setState({password: event.target.value});
        break;
    }
  }

  handleSubmit() {
    authenticationService.login(this.state);
  }

  render() {
    const { classes } = this.props;

    return (
      <form className={classes.container} autoComplete="off">
        <TextField
          className={classes.textField}
          id="standard-name"
          label="Name"
          margin="normal"
          name="name"
          onChange={this.handleChange}
          value={this.state.name}
          required
        />

        <TextField
          className={classes.textField}
          id="standard-password-input"
          label="Password"
          margin="normal"
          name="password"
          onChange={this.handleChange}
          type="password"
          value={this.state.password}
          required
        />

        <Button
          className={classes.button}
          color="primary"
          onClick={this.handleSubmit}
          type="button"
          variant="contained">
          Login
        </Button>
      </form>
    );
  }
}

Login.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Login);