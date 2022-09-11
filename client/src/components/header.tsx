import { Container, Nav, Navbar } from 'react-bootstrap'

const Header = () => {
  return (
     <Navbar bg='dark' variant='dark' expand='lg' collapseOnSelect>
       <Container>
         <Navbar.Brand href='/'>Go React Auth</Navbar.Brand>
         <Navbar.Toggle aria-controls='basic-navbar-nav'/>
         <Navbar.Collapse id='basic-navbar-nav'>
           <Nav className='ms-auto'>
             <Nav.Link href='/registration'>Registration</Nav.Link>
             <Nav.Link href='/login'>Login</Nav.Link>
           </Nav>
         </Navbar.Collapse>
       </Container>
     </Navbar>
  )
}

export default Header