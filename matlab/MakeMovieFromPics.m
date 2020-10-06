clear all;

writerObj = VideoWriter('d:\MatlabMovie3\Y1movie2.avi');
writerObj.FrameRate = 4;
open(writerObj);

for t=1:51
frame = im2frame(imread(['d:\MatlabMovie3\New\pics Y1\',num2str(t),'.tif']));
writeVideo(writerObj,frame);
end;

for t=1:51
frame = im2frame(imread(['d:\MatlabMovie3\New\pics Y3\',num2str(t),'.tif']));
writeVideo(writerObj,frame);
end;


close(writerObj);