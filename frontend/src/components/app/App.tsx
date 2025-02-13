import { BrowserRouter } from 'react-router-dom';
import AppRouter from './AppRouter';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <AppRouter />;
    </BrowserRouter>
  );
};

export default App;
