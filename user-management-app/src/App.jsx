import { Container, CssBaseline, ThemeProvider, createTheme } from '@mui/material';
import UserForm from './components/UserForm';
import UserList from './components/UserList';

const theme = createTheme();

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Container maxWidth="md" sx={{ py: 4 }}>
        <UserForm />
        <UserList />
      </Container>
    </ThemeProvider>
  );
}

export default App;
