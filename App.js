import { StoreProvider } from "@/store";
import { QueryClientProvider } from "@tanstack/react-query";
import HomeLayout from "./src/app/index";
import { client } from "./src/shared/query";

const App = () => {
  return (
    <QueryClientProvider client={client}>
      <StoreProvider>
        <HomeLayout />
      </StoreProvider>
    </QueryClientProvider>
  );
};

export default App;
