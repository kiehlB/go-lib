import { PageGrid } from "@/components/grid";
import Navbar from "@/components/navbar/navbar-item";

export default function MainLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <PageGrid as='div' className='pb-[2rem] pt-[2rem]'>
      <aside className='col-span-2'>
        <Navbar
          primaryItems={[
            {
              text: "Home",
              to: "/",
              sub: ["/search", "/search/[query]", "/trending"],
            },
            {
              text: "Today",
              to: "/today",
            },
            {
              text: "Done",
              to: "/tags",
            },
          ]}
          secondaryItems={[
            {
              text: "Trending tags",
              to: "/Trending tags",
            },
          ]}
        />
      </aside>

      {children}
      <div className='pt-4 col-span-3'>last</div>
    </PageGrid>
  );
}
