import { render, screen } from '@testing-library/react';
import App from '../App';

// Mock Ant Design components
jest.mock('antd', () => ({
  ...jest.requireActual('antd'),
  Layout: {
    Header: ({ children }: { children: React.ReactNode }) => <header data-testid="header">{children}</header>,
    Content: ({ children }: { children: React.ReactNode }) => <main data-testid="content">{children}</main>,
    Footer: ({ children }: { children: React.ReactNode }) => <footer data-testid="footer">{children}</footer>,
  },
  Typography: {
    Title: ({ children }: { children: React.ReactNode }) => <h1>{children}</h1>,
  },
  Card: ({ children }: { children: React.ReactNode }) => <div data-testid="card">{children}</div>,
  Statistic: ({ title, value }: { title: string; value: string | number }) => (
    <div data-testid="statistic">
      <span>{title}</span>
      <span>{value}</span>
    </div>
  ),
  Row: ({ children }: { children: React.ReactNode }) => <div data-testid="row">{children}</div>,
  Col: ({ children }: { children: React.ReactNode }) => <div data-testid="col">{children}</div>,
}));

describe('App Component', () => {
  it('renders without crashing', () => {
    render(<App />);
    expect(screen.getByText(/Blockchain CMDB/i)).toBeInTheDocument();
  });

  it('displays dashboard statistics', () => {
    render(<App />);
    expect(screen.getByText(/Total Assets/i)).toBeInTheDocument();
    expect(screen.getByText(/Blockchain Records/i)).toBeInTheDocument();
    expect(screen.getByText(/Security Status/i)).toBeInTheDocument();
  });

  it('displays welcome message', () => {
    render(<App />);
    expect(screen.getByText(/Welcome to Blockchain CMDB/i)).toBeInTheDocument();
  });

  it('displays feature list', () => {
    render(<App />);
    expect(screen.getByText(/Asset management with blockchain audit trail/i)).toBeInTheDocument();
    expect(screen.getByText(/Real-time dashboard and analytics/i)).toBeInTheDocument();
  });

  it('displays footer', () => {
    render(<App />);
    expect(screen.getByText(/Created by OpenClaw Agent/i)).toBeInTheDocument();
  });
});
